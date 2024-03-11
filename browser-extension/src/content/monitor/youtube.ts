import { Monitor } from './abstract'

export class YouTube extends Monitor {
    name = 'youtube'
    private readonly element: HTMLVideoElement

    externalPause = false
    externalPlay = false

    constructor (port: chrome.runtime.Port) {
        super(port)
        console.log(`Kimparty: ${this.name} monitor initialized`)

        const element = document.querySelector('video')

        if (element === null) {
            throw new Error('Kimparty: YouTube video element not found')
        }

        this.element = element

        this.setupPlayListener()
        this.setupPauseListener()
        element.pause()
    }

    setupPlayListener (): void {
        this.element.addEventListener('play', () => {
            if (this.externalPlay) {
                this.externalPlay = false
                return
            }

            const content = `play:${this.element.currentTime.toString()}`
            this.sendPortMessage(content)
        })
    }

    setupPauseListener (): void {
        this.element.addEventListener('pause', () => {
            if (this.externalPause) {
                this.externalPause = false
                return
            }

            const content = `pause:${this.element.currentTime.toString()}`
            this.sendPortMessage(content)
        })
    }

    setupSeekedListener (): void {
        this.element.addEventListener('seeked', () => {
            const content = `sync:${this.element.currentTime.toString()}`
            this.sendPortMessage(content)
        })
    }

    play (time: number): void {
        this.externalPlay = true
        this.sync(time)
        this.element.play().catch(console.error)
    }

    pause (time: number): void {
        this.externalPause = true
        this.element.pause()
        this.sync(time)
    }

    sync (time: number): void {
        if (time < 0) {
            return
        }

        this.element.currentTime = time
    }
}
