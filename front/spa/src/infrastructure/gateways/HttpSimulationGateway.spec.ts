import { describe, expect, it, vi } from "vitest";
import { HttpSimulationGateway } from "./HttpSimulationGateway";
import { readBaseUrl } from "./readBaseUrl";
import { requestJson } from "./requestJson";

vi.mock("./readBaseUrl", () => ({
    readBaseUrl: vi.fn(() => "http://example.test"),
}));

vi.mock("./requestJson", () => ({
    requestJson: vi.fn(),
}));

describe("HttpSimulationGateway", () => {
    it("uses provided base url and delegates to requestJson", async () => {
        const gateway = new HttpSimulationGateway("http://custom.test");
        const requestJsonMock = requestJson as unknown as ReturnType<typeof vi.fn>;
        requestJsonMock.mockResolvedValue({ data: true });

        await gateway.fetchBootstrap();
        await gateway.simulate({ teams: [], weeks: [] });

        expect(requestJson).toHaveBeenCalledWith("http://custom.test/bootstrap");
        expect(requestJson).toHaveBeenCalledWith(
            "http://custom.test/simulate",
            expect.objectContaining({ method: "POST" }),
        );
    });

    it("throws when base url is missing", () => {
        const readerMock = readBaseUrl as unknown as ReturnType<typeof vi.fn>;
        readerMock.mockImplementation(() => {
            throw new Error("missing");
        });

        expect(() => new HttpSimulationGateway()).toThrowError();
    });
});
