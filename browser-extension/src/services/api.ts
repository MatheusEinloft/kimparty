const fromEnv = import.meta.env.API_BASE_URL
export const API_BASE_URL = fromEnv !== undefined ? fromEnv : 'http://127.0.0.1:3000'

interface PartyData {
    id: string
    url: string
}

export async function createParty (videoURL: string): Promise<PartyData> {
    try {
        const res = await fetch(`${API_BASE_URL}/party?url=${videoURL}`, {
            method: 'POST'
        })

        const data = await res.json()
        return { id: data.id, url: data.url }
    } catch (error) {
        throw new Error('Kimparty: failed to create party')
    }
}

export async function findParty (partyID: string): Promise<PartyData> {
    try {
        const res = await fetch(`${API_BASE_URL}/party/find?id=${partyID}`)
        const data = await res.json()
        return { id: data.id, url: data.url }
    } catch (error) {
        throw new Error('Kimparty: failed to find party')
    }
}
