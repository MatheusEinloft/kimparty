import { type ChromeMessage } from '../services/chrome'
import { type GlobalStoreFields } from '../store/global'
import { type Monitor } from './monitor/abstract'
import { YouTube } from './monitor/youtube'
import { getPersistentStore, persistentStore, setPersistentStore } from './persistent'

console.log('Kimparty: Content script loaded')
window.addEventListener('popstate', cleanup)

let monitor: Monitor | null = null
let port: chrome.runtime.Port | null = null

chrome.runtime.onMessage.addListener(listener)

function listener(message: ChromeMessage, _: chrome.runtime.MessageSender, sendResponse: (reponse?: any) => void): void {
    switch (message.author) {
        case 'popup':
            handleFromPopup(message, sendResponse)
            break
    }
}

function handleFromPopup(message: ChromeMessage, sendRespone: (reponse?: any) => void): void {
    switch (message.action) {
        case 'set-store':
            setPersistentStore(message.payload as GlobalStoreFields)
            break
        case 'get-store':
            sendRespone(getPersistentStore())
            break
        case 'start':
            start()
            break
        case 'stop':
            cleanup()
            break
        case 'ping':
            sendRespone('pong')
            break
    }
}

function start(): void {
    if (persistentStore.PARTY_ID === '') {
        console.log('Kimparty: Party ID not set')
        return
    }

    if (monitor !== null) {
        console.log('Kimparty: Monitor already running')
        return
    }

    port = chrome.runtime.connect({ name: 'worker' })

    port.postMessage({
        action: 'start', payload: generateWebScoketURL(), author: 'monitor'
    })

    monitor = new YouTube(port)

}

function cleanup(): void {
    console.log('Kimparty: Cleaning up content script...')

    monitor?.cleanup()
    monitor = null

    port?.postMessage({ action: 'stop', author: 'monitor', payload: null })
    port?.disconnect()

    chrome.runtime.onMessage.removeListener(listener)
}

function generateWebScoketURL(): string {
    const apiUrl = persistentStore.API_URL.replace('http', 'ws').replace('https', 'wss')
    return `${apiUrl}/ws/join?party_id=${persistentStore.PARTY_ID}&username=`
}
