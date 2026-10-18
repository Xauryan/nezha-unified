import { type PersistenceSettingsContent } from "@cert/domain/settings";

export interface PersistenceSettingsState {
  settings: PersistenceSettingsContent;
  loading: boolean;
  loadedAtOnce: boolean;
}

export interface PersistenceSettingsActions {
  loadSettings: (refresh?: boolean) => Promise<void>;
  saveSettings: (settings: PersistenceSettingsContent) => Promise<void>;
}

export type PersistenceSettingsStore = PersistenceSettingsState & PersistenceSettingsActions;
