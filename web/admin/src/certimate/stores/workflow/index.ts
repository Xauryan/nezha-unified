import { registerSessionReset } from "@cert/session";
import { produce } from "immer";
import { isEqual } from "radash";
import { create } from "zustand";

import { WORKFLOW_NODE_TYPES, type WorkflowModel, type WorkflowNodeConfigForStart } from "@cert/domain/workflow";
import { get as getWorkflow, save as saveWorkflow, subscribe as subscribeWorkflow } from "@cert/repository/workflow";

import { type WorkflowStore } from "./types";

export const useWorkflowStore = create<WorkflowStore>((set, get) => {
  const ensureInitialized = () => {
    if (!get().initialized) throw "Workflow not initialized yet";
  };

  let unsubscriber: (() => void) | undefined;
  let lifecycle = 0;

  registerSessionReset(() => get().destroy());

  return {
    workflow: {} as WorkflowModel,
    initialized: false,

    init: async (id: string) => {
      const current = ++lifecycle;
      unsubscriber?.();
      unsubscriber = undefined;
      set({ workflow: {} as WorkflowModel, initialized: false });
      const data = await getWorkflow(id);
      if (current !== lifecycle) return;
      set({
        workflow: data,
        initialized: true,
      });

      const stop = await subscribeWorkflow(id, (cb) => {
        if (current !== lifecycle || cb.record.id !== get().workflow.id) return;

        set({
          workflow: cb.record,
        });
      });
      if (current === lifecycle) unsubscriber = stop;
      else stop();
    },

    destroy: () => {
      lifecycle++;
      unsubscriber?.();
      unsubscriber = void 0;

      set({
        workflow: {} as WorkflowModel,
        initialized: false,
      });
    },

    setName: async (name) => {
      ensureInitialized();
      const current = lifecycle;

      const resp = await saveWorkflow({
        id: get().workflow.id!,
        name: name || "",
      });

      if (current !== lifecycle) return;
      set((state) => {
        return {
          workflow: produce(state.workflow, (draft) => {
            draft.name = resp.name;
          }),
        };
      });
    },

    setDescription: async (description) => {
      ensureInitialized();
      const current = lifecycle;

      const resp = await saveWorkflow({
        id: get().workflow.id!,
        description: description || "",
      });

      if (current !== lifecycle) return;
      set((state) => {
        return {
          workflow: produce(state.workflow, (draft) => {
            draft.description = resp.description;
          }),
        };
      });
    },

    setEnabled: async (enabled) => {
      ensureInitialized();
      const current = lifecycle;

      const resp = await saveWorkflow({
        id: get().workflow.id!,
        enabled: enabled,
      });

      if (current !== lifecycle) return;
      set((state) => {
        return {
          workflow: produce(state.workflow, (draft) => {
            draft.enabled = resp.enabled;
          }),
        };
      });
    },

    orchestrate: async (graph) => {
      ensureInitialized();
      const current = lifecycle;

      const resp = await saveWorkflow({
        id: get().workflow.id!,
        graphDraft: graph,
        hasDraft: !isEqual(graph, get().workflow.graphContent),
      });

      if (current !== lifecycle) return;
      set((state) => {
        return {
          workflow: produce(state.workflow, (draft) => {
            draft.graphDraft = resp.graphDraft;
            draft.hasDraft = resp.hasDraft;
          }),
        };
      });
    },

    publish: async () => {
      ensureInitialized();
      const current = lifecycle;

      const graph = get().workflow.graphDraft!;
      if (graph?.nodes?.[0]?.type !== WORKFLOW_NODE_TYPES.START) throw "Workflow nodes tree of draft in invalid";
      const startConfig = graph.nodes[0].data.config as WorkflowNodeConfigForStart;
      const resp = await saveWorkflow({
        id: get().workflow.id!,
        trigger: startConfig.trigger,
        triggerCron: startConfig.triggerCron,
        graphContent: graph,
        hasContent: true,
        hasDraft: false,
      });

      if (current !== lifecycle) return;
      set((state) => {
        return {
          workflow: produce(state.workflow, (draft) => {
            draft.trigger = resp.trigger;
            draft.triggerCron = resp.triggerCron;
            draft.graphContent = resp.graphContent;
            draft.hasContent = resp.hasContent;
            draft.hasDraft = resp.hasDraft;
          }),
        };
      });
    },

    rollback: async () => {
      ensureInitialized();
      const current = lifecycle;

      const graph = get().workflow.graphContent!;
      if (graph?.nodes?.[0]?.type !== WORKFLOW_NODE_TYPES.START) throw "Workflow nodes tree of content in invalid";
      const startConfig = graph.nodes[0].data.config as WorkflowNodeConfigForStart;
      const resp = await saveWorkflow({
        id: get().workflow.id!,
        trigger: startConfig.trigger,
        triggerCron: startConfig.triggerCron,
        hasContent: true,
        graphDraft: graph,
        hasDraft: false,
      });

      if (current !== lifecycle) return;
      set((state) => {
        return {
          workflow: produce(state.workflow, (draft) => {
            draft.trigger = resp.trigger;
            draft.triggerCron = resp.triggerCron;
            draft.hasContent = resp.hasContent;
            draft.graphDraft = resp.graphDraft;
            draft.hasDraft = resp.hasDraft;
          }),
        };
      });
    },
  };
});
