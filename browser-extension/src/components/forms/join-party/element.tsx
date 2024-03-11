import { createSignal, type Component, Show } from 'solid-js'
import Input from '../../input/input'
import Button from '../../buttons/button'
import styles from './styles.module.css'
import { validatePartyID } from '../../../services/kimparty'
import { findParty } from '../../../services/api'
import { pingContentScript, redirectTab, sendMessageToContent } from '../../../services/chrome'
import { GlobalStore, PersistGlobalStore, SetGlobalStore, State } from '../../../store/global'

const handleContentScriptLoaded = async (): Promise<void> => {
    console.log('Kimparty: waiting for content script to load...')

    await new Promise(resolve => setTimeout(resolve, 2000))
    let contentLoaded = await pingContentScript()

    let attempts = 0
    while (!contentLoaded && attempts < 10) {
        await new Promise(resolve => setTimeout(resolve, 2000))

        console.log('Kimparty: content script not loaded, retrying...')
        contentLoaded = await pingContentScript()

        if (contentLoaded) {
            console.log('Kimparty: content script loaded')
            break
        }

        attempts++
    }

    if (!contentLoaded) {
        throw new Error('Kimparty: content script not loaded')
    }

    console.log('Kimparty: content script loaded')
}

export const FormJoinParty: Component = () => {
    const [partyID, setPartyID] = createSignal('')
    const [valid, setValid] = createSignal({
        valid: true,
        message: ''
    })

    const handleOnSubmit = async (e: Event): Promise<void> => {
        e.preventDefault()

        if (GlobalStore.State === State.InParty) {
            return
        }

        const isValid = validatePartyID(partyID())

        if (!isValid) {
            setPartyID(partyID())
            setValid({
                valid: false,
                message: 'INVALID ID'
            })
            return
        }

        try {
            const foundParty = await findParty(partyID())
            SetGlobalStore({
                PARTY_ID: foundParty.id,
                State: State.InParty
            })

            console.log('Kimparty: redirecting to', foundParty.url)
            await redirectTab(foundParty.url)

            await handleContentScriptLoaded()
            await PersistGlobalStore()
            await sendMessageToContent({ action: 'start', author: 'popup', payload: null })
        } catch (error: any) {
            console.error('Kimparty: error joining party', error)
            setValid({
                valid: false,
                message: error as string
            })
        }
    }

    return (
        <form class={styles.form} onSubmit={(e) => { handleOnSubmit(e).catch(console.error) }}>
            <Show when={!valid().valid}>
                <span class={styles.error}>{valid().message}</span>
            </Show>
            <Input placeholder='Paste ID' value={GlobalStore.State === State.InParty ? GlobalStore.PARTY_ID : partyID()} setVal={setPartyID} />
            <Button value={GlobalStore.State === State.InParty ? 'in party' : 'Join a Party'} />
        </form>
    )
}
