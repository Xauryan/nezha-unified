import { type WorkflowGraph, type WorkflowModel } from "@cert/domain/workflow";

export interface WorkflowState {
  workflow: WorkflowModel;
  initialized: boolean;
}

export interface WorkflowActions {
  init(id: string): Promise<void>;
  destroy(): void;

  setName: (name: Required<WorkflowModel>["name"]) => Promise<void>;
  setDescription: (description: Required<WorkflowModel>["description"]) => Promise<void>;
  setEnabled(enabled: Required<WorkflowModel>["enabled"]): Promise<void>;

  orchestrate(graph: WorkflowGraph): Promise<void>;
  publish(): Promise<void>;
  rollback(): Promise<void>;
}

export interface WorkflowStore extends WorkflowState, WorkflowActions {}
