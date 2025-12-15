import type { SimulationRequestPayload } from "../../domain/entities/SimulationRequestPayload";
import type { SimulationResponse } from "../../domain/entities/SimulationResponse";
import type { SimulationGateway } from "../../domain/ports/SimulationGateway";
import { readBaseUrl } from "./readBaseUrl";
import { requestJson } from "./requestJson";

export class HttpSimulationGateway implements SimulationGateway {
    private readonly baseUrl: string;

    constructor(baseUrl = readBaseUrl()) {
        this.baseUrl = baseUrl;
    }

    async fetchBootstrap(): Promise<SimulationResponse> {
        return requestJson<SimulationResponse>(`${this.baseUrl}/bootstrap`);
    }

    async simulate(payload: SimulationRequestPayload): Promise<SimulationResponse> {
        return requestJson<SimulationResponse>(`${this.baseUrl}/simulate`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(payload),
        });
    }
}
