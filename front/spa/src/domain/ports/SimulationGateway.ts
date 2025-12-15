import type { SimulationRequestPayload } from "../entities/SimulationRequestPayload";
import type { SimulationResponse } from "../entities/SimulationResponse";

export interface SimulationGateway {
    fetchBootstrap(): Promise<SimulationResponse>;
    simulate(payload: SimulationRequestPayload): Promise<SimulationResponse>;
}
