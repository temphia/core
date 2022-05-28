export const PAGE_AUTH_MAIN = "main_auth";
export const PAGE_AFTER_AUTH = "after_auth";
export const PAGE_EXTERNAL_AUTH = "external_auth";

export const parseParams = async () => {
  return {
    page_type: PAGE_AUTH_MAIN,
    data: {
      group: "",
      tenant_id: "",
      base_url: "",
    },
  };
};
