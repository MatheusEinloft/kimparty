import { createEffect, type Component } from 'solid-js'
import { GlobalStore, PersistGlobalStore, SetGlobalStore, State } from '../../store/global'
import { create } from '../../services/kimparty'
import { sendMessageToContent } from '../../services/chrome'

const onClick = async (): Promise<void> => {
    switch (GlobalStore.State) {
        case State.Creatable:
            await create()
            await PersistGlobalStore()
            await sendMessageToContent({ action: 'start', author: 'popup', payload: null })
            break
        case State.CreatedParty:
            await navigator.clipboard.writeText(GlobalStore.PARTY_ID)
            SetGlobalStore('State', State.CopiedPartyID)
            break
    }
}

const value = (): string => {
    switch (GlobalStore.State) {
        case State.CreatingParty:
            return 'Creating Party...'
        case State.CreatedParty:
            return 'Copy to clipboard'
        case State.CopiedPartyID:
            return 'Copied!'
        default:
            return 'Create Party'
    }
}

export const ButtonCreateParty: Component = () => {
    createEffect(() => {
        if (GlobalStore.State === State.CopiedPartyID) {
            setTimeout(() => {
                SetGlobalStore('State', State.CreatedParty)
            }, 800)
        }
    })

    return (
        // eslint-disable-next-line @typescript-eslint/no-misused-promises
        <button onClick={onClick} class='button button-primary'>
            {value()}
        </button>
    )
}
