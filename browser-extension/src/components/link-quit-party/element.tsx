import type { Component } from 'solid-js'
import styles from './styles.module.css'
import { sendMessageToContent } from '../../services/chrome'
import { PersistGlobalStore, SetGlobalStore, State } from '../../store/global'

const LinkQuitParty: Component = () => {
    const handleQuit = async (e: Event): Promise<void> => {
        e.preventDefault()
        console.log('Kimparty: Quitting party...')

        await sendMessageToContent({ action: 'stop', author: 'popup', payload: null })

        SetGlobalStore({ State: State.Creatable, PARTY_ID: '' })
        await PersistGlobalStore()
    }

    return (
        <a onClick={(e) => handleQuit(e).catch(console.error)} class={styles.quit} href='#' >Quit</a>
    )
}

export default LinkQuitParty
