import { describe, expect, it, vi } from "vitest";
import { readBaseUrl } from "./readBaseUrl";

describe("readBaseUrl", () => {
    it("trims trailing slashes", () => {
        vi.stubEnv("VITE_BFF_BASE_URL", "http://example.test/");
        expect(readBaseUrl()).toBe("http://example.test");
        vi.unstubAllEnvs();
    });

    it("throws when missing", () => {
        vi.stubEnv("VITE_BFF_BASE_URL", "");
        expect(() => readBaseUrl()).toThrowError();
        vi.unstubAllEnvs();
    });
});
