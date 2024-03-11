import { GlobalStore, SetGlobalStore, State } from '../store/global'
import { createParty } from './api'

export function validatePartyID (partyID: string): boolean {
    const uuidRegex = /^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i
    return uuidRegex.test(partyID)
}

export async function create (): Promise<void> {
    const partyData = await createParty(GlobalStore.URL)

    if (partyData === null) {
        return
    }

    SetGlobalStore({ PARTY_ID: partyData.id, State: State.CreatedParty })
}

export function generateURLWithQuery (url: string, paramValue: string): string {
    const separator = url.includes('?') ? '&' : '?'
    return `${url}${separator}kimparty_id=${encodeURIComponent(paramValue)}`
}
