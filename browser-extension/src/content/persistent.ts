import { API_BASE_URL } from '../services/api'
import { State, type GlobalStoreFields } from '../store/global'

export const persistentStore = {
    State: State.Creatable,
    URL: window.location.href,
    API_URL: API_BASE_URL,
    PARTY_ID: ''
}

export function getPersistentStore (): typeof persistentStore {
    return persistentStore
}

export function setPersistentStore (store: GlobalStoreFields): void {
    Object.assign(persistentStore, store)
}
