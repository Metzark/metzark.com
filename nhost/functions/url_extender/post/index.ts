import type { Request, Response } from "express";
import { randomBytes } from "crypto";

const MAX_EXTENDED_URL_LENGTH = 4096;

const insertUrlGql = `mutation InsertUrl($actual: String!, $extended: String!, $expires_at: timestamp, $required_check: Boolean) {
  insert_url_extender_urls_one(object: {actual: $actual, extended: $extended, expires_at: $expires_at, require_check: $required_check}) {
    actual
    extended
    expires_at
    require_check
    created_at
  }
}`;

export default async (req: Request, res: Response) => {
  try {
    const { actual, expires_at, required_check, length } = req.body;

    if (!actual) {
      return res.status(400).json({
        error: "actual is required",
      });
    }

    if (length > MAX_EXTENDED_URL_LENGTH) {
      return res.status(400).json({
        error: `extended url is too long. keep the length under ${MAX_EXTENDED_URL_LENGTH} characters`,
      });
    }

    const extended = generateLongString(length);

    const expiresAt = expires_at ?? new Date(Date.now() + 24 * 60 * 60 * 1000).toISOString();

    const response = await fetch(process.env.NHOST_GRAPHQL_URL!, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "x-hasura-admin-secret": process.env.NHOST_ADMIN_SECRET!,
      },
      body: JSON.stringify({
        query: insertUrlGql,
        variables: {
          actual,
          extended,
          expires_at: expiresAt,
          required_check: required_check ?? false,
        },
      }),
    });

    const result = await response.json();

    if (!response.ok || result.errors) {
      console.error(result);

      return res.status(500).json({
        error: "Failed to insert URL",
      });
    }

    return res.status(201).json(result.data.insert_url_extender_urls_one);
  } catch (error) {
    console.error(error);

    return res.status(500).json({
      error: "Internal server error",
    });
  }
};

function generateLongString(length = 2048): string {
  const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";

  let result = "";

  while (result.length < length) {
    const bytes = randomBytes(length);

    for (const byte of bytes) {
      result += chars[byte % chars.length];

      if (result.length === length) {
        break;
      }
    }
  }

  return result;
}
