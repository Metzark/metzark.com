import type { Request, Response } from "express";

const getUrlGql = `query GetUrl($extended: String!) {
  url_extender_urls_by_pk(extended: $extended) {
    actual
    extended
    expires_at
    require_check
    created_at
  }
}`;

export default async (req: Request, res: Response) => {
  try {
    const { extended } = req.query;
    console.log({ extended });

    if (!extended) {
      return res.status(400).json({
        error: "extended is required",
      });
    }

    const response = await fetch(process.env.NHOST_GRAPHQL_URL!, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "x-hasura-admin-secret": process.env.NHOST_ADMIN_SECRET!,
      },
      body: JSON.stringify({
        query: getUrlGql,
        variables: {
          extended,
        },
      }),
    });

    const result = await response.json();

    if (!response.ok || result.errors) {
      console.error(result);

      return res.status(500).json({
        error: "Failed to retrieve URL",
      });
    }

    const url = result.data.url_extender_urls_by_pk;

    if (!url) {
      return res.status(404).json({
        error: "URL not found",
      });
    }

    return res.status(200).json(url);
  } catch (error) {
    console.error(error);

    return res.status(500).json({
      error: "Internal server error",
    });
  }
};
