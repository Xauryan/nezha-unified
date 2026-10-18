import { type SSLProviderSettingsContent } from "@cert/domain/settings";

export interface SSLProviderSettingsState {
  settings: SSLProviderSettingsContent;
  loading: boolean;
  loadedAtOnce: boolean;
}

export interface SSLProviderSettingsActions {
  loadSettings: (refresh?: boolean) => Promise<void>;
  saveSettings: (settings: SSLProviderSettingsContent) => Promise<void>;
}

export type SSLProviderSettingsStore = SSLProviderSettingsState & SSLProviderSettingsActions;
