import { createStore } from 'solid-js/store'
import { API_BASE_URL } from '../services/api'
import { sendMessageToContent } from '../services/chrome'

export enum State {
    Default,
    Creatable,
    CreatingParty,
    CreatedParty,
    CopiedPartyID,
    JoiningParty,
    InParty,
}

export interface GlobalStoreFields {
    State: State
    URL: string
    PARTY_ID: string
    API_URL: string
}

export const [GlobalStore, SetGlobalStore] = createStore<GlobalStoreFields>({
    State: State.Default,
    URL: '',
    PARTY_ID: '',
    API_URL: API_BASE_URL
})

export async function PersistGlobalStore (): Promise<void> {
    await sendMessageToContent({
        action: 'set-store',
        payload: GlobalStore,
        author: 'popup'
    })
}

export async function LoadGlobalStore (): Promise<boolean> {
    try {
        await sendMessageToContent({
            action: 'get-store',
            payload: null,
            author: 'popup'
        }, (response: GlobalStoreFields) => {
            if (response === undefined) {
                return false
            }

            SetGlobalStore(response)
            return true
        })
    } catch (e) {
        console.error(e)
    }

    return false
}
