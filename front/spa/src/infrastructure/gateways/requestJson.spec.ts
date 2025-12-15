import { describe, expect, it, vi } from "vitest";
import { requestJson } from "./requestJson";

describe("requestJson", () => {
    it("throws on non-ok responses", async () => {
        const mockResponse = new Response("", { status: 500 });
        vi.stubGlobal("fetch", vi.fn().mockResolvedValue(mockResponse));

        await expect(requestJson("/test")).rejects.toThrowError("Request failed");

        vi.unstubAllGlobals();
    });
});
