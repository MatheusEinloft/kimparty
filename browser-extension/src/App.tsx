import { type Component, Show, createEffect } from 'solid-js'
import { GlobalStore, LoadGlobalStore, SetGlobalStore, State } from './store/global'

import Layout from './layouts/layout'
import Button from './components/buttons/button'
import { ButtonCreateParty } from './components/buttons/create-party'
import { FormJoinParty } from './components/forms/join-party/element'
import Input from './components/input/input'
import LinkQuitParty from './components/link-quit-party/element'
import { getCurrentValidURL, injectContentScript, pingContentScript } from './services/chrome'

// Create
const showCreatePartyButton = (): boolean => {
    return GlobalStore.State !== State.JoiningParty &&
        GlobalStore.State !== State.Default &&
        GlobalStore.State !== State.InParty
}

const showCreatePartyInput = (): boolean => {
    return GlobalStore.State !== State.Default &&
        GlobalStore.State !== State.JoiningParty &&
        GlobalStore.State !== State.Creatable &&
        GlobalStore.State !== State.InParty
}

// Join
const showJoinPartyButton = (): boolean => {
    return GlobalStore.State === State.Creatable
}

const showJoinPartyForm = (): boolean => {
    return GlobalStore.State === State.Default ||
        GlobalStore.State === State.JoiningParty ||
        GlobalStore.State === State.InParty
}

const onPopupOpen = async (): Promise<void> => {
    const url = await getCurrentValidURL()

    if (url === null) {
        return
    }

    console.log('Kimparty: Valid URL', url)

    console.log('Kimparty: Checking if content script is loaded...')
    const contentLoaded = await pingContentScript()

    if (!contentLoaded) {
        console.log('Kimparty: Content script not loaded, injecting...')
        await injectContentScript()
    }

    const loaded = await LoadGlobalStore()

    if (loaded) {
        return
    }

    SetGlobalStore({
        State: State.Creatable,
        URL: url
    })
}

const App: Component = () => {
    createEffect(async () => {
        await onPopupOpen()
    })

    return (
        <Layout title='Kimparty' subtitle='Watch movies with your friends'>
            <Show when={GlobalStore.State === State.InParty}>
                <LinkQuitParty />
            </Show>
            <Show when={showCreatePartyInput()}>
                <Input value={GlobalStore.PARTY_ID !== '' ? GlobalStore.PARTY_ID : 'Loading...'} readonly={true} />
            </Show>
            <Show when={showCreatePartyButton()}>
                <ButtonCreateParty />
            </Show>

            <Show when={showJoinPartyForm()}>
                <FormJoinParty />
            </Show>
            <Show when={showJoinPartyButton()}>
                <Button onClick={() => { SetGlobalStore('State', State.JoiningParty as State) }} value="Join a Party" />
            </Show>
        </Layout>
    )
}

export default App
