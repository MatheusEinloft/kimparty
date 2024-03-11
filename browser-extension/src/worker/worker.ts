import { type ChromeMessage } from '../services/chrome'
import { newWebSocket } from './websocket'

let ws: WebSocket | null = null

chrome.runtime.onConnect.addListener((port) => {
    if (port.name !== 'worker') {
        return
    }

    port.onMessage.addListener((message: ChromeMessage) => {
        if (message.author !== 'monitor') {
            return
        }

        switch (message.action) {
            case 'stop':
                cleanup()
                break
            case 'start':
                ws = newWebSocket(message.payload as string, port)
                break
            case 'message':
                ws?.send(JSON.stringify(message.payload))
                break
        }
    })

    port.onDisconnect.addListener(() => {
        console.log('Kimparty: Worker port disconnected')
        cleanup()
    })
})

function cleanup(): void {
    console.log('Kimparty: Cleaning up worker...')

    if (ws !== null) {
        ws.close()
    }

    ws = null
}
