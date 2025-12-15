export function readBaseUrl(raw = import.meta.env.VITE_BFF_BASE_URL ?? ""): string {
    const trimmed = raw.replace(/\/$/, "");
    if (!trimmed) {
        throw new Error("VITE_BFF_BASE_URL environment variable is required");
    }
    return trimmed;
}
