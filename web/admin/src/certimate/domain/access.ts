export interface AccessModel extends BaseModel {
  name: string;
  provider: string;
  config?: Record<string, unknown>;
  reserve?: "ca" | "notif";
  unified_version?: number;
  unified_enabled?: boolean;
  configConfigured?: boolean;
}
