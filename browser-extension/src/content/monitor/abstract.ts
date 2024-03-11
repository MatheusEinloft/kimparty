import { type ChromeMessage } from '../../services/chrome'
import { type WebSocketMessage } from '../../worker/websocket'

export abstract class Monitor {
    protected port: chrome.runtime.Port
    abstract externalPause: boolean
    abstract externalPlay: boolean
    abstract name: string

    constructor (port: chrome.runtime.Port) {
        this.port = port


        this.port.onMessage.addListener((message: ChromeMessage) => {
            if (message.action !== 'message' || message.author !== 'websocket') {
                return
            }

            this.onMessage(message.payload as WebSocketMessage)
        })
    }

    protected onMessage (message: WebSocketMessage): void {
        if (message.type !== 1) {
            return
        }

        let [content, time] = message.content.split(':')

        if (time === undefined) {
            time = '-1'
        }

        switch (content) {
            case 'play':
                this.play(Number(time))
                break
            case 'pause':
                this.pause(Number(time))
                break
            case 'sync':
                this.sync(Number(time))
                break
        }
    }

    protected pingPort (): void {
        this.port.postMessage({ action: 'message', author: 'monitor', payload: 'ping' })
    }

    protected sendPortMessage (content: string, type: number = 1): void {
        this.port.postMessage({ action: 'message', author: 'monitor', payload: { content, type } })
    }

    abstract sync (time: number): void
    abstract play (time: number): void
    abstract pause (time: number): void

    abstract setupPlayListener (): void
    abstract setupPauseListener (): void
    abstract setupSeekedListener (): void

    cleanup (): void {
        this.port.disconnect()
    }
}
