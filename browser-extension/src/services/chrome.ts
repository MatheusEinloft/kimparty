const STREAMING_VIDEO_PLATFORMS = [
    'youtube.com/watch?v=',
    'netflix.com/watch'
]

export interface ChromeMessage {
    action: string
    payload: any
    author: string
}

export async function sendMessageToContent (message: ChromeMessage, callback?: (response?: any) => void): Promise<void> {
    try {
        const tabID = await getCurrentTabID()

        if (tabID === -1) {
            return
        }

        if (callback !== undefined) {
            chrome.tabs.sendMessage(tabID, message, callback)
            return
        }

        await chrome.tabs.sendMessage(tabID, message)
    } catch (error) {
        console.error(error)
    }
}

export async function sendMessageToWorker (message: ChromeMessage): Promise<void> {
    try {
        await chrome.runtime.sendMessage(message)
    } catch (error) {
        console.error(error)
    }
}

export async function getCurrentTabID (): Promise<number> {
    try {
        const tabs = await chrome.tabs.query({ active: true, currentWindow: true })
        const tab = tabs[0]

        if (tab.id === undefined) {
            return -1
        }

        return tab.id
    } catch (error) {
        console.error(error)
        return -1
    }
}

export async function pingContentScript (): Promise<boolean> {
    try {
        const tabID = await getCurrentTabID()

        if (tabID === -1) {
            return false
        }

        return await new Promise(resolve => {
            chrome.tabs.sendMessage(tabID, { action: 'ping', author: 'popup' }, (response) => {
                resolve(response === 'pong')
            })
        })
    } catch (error) {
        console.error(error)
        return false
    }
}

export async function injectContentScript (): Promise<void> {
    try {
        const tabID = await getCurrentTabID()

        if (tabID === -1) {
            return
        }

        const manifest = chrome.runtime.getManifest()

        if (manifest?.content_scripts === undefined) {
            return
        }

        await Promise.all(
            manifest.content_scripts.map(async (script): Promise<void> => {
                console.log(script)
                if (script.js === undefined) {
                    return
                }

                for (const js of script.js) {
                    await chrome.scripting.executeScript({
                        target: { tabId: tabID },
                        files: [js]
                    })
                }
            })
        )
    } catch (error) {
        console.error(error)
    }
}

export async function redirectTab (url: string): Promise<void> {
    try {
        const tabID = await getCurrentTabID()

        if (tabID === -1) {
            return
        }

        await chrome.tabs.update(tabID, { url })
    } catch (error) {
        console.error(error)
    }
}

export async function getCurrentValidURL (): Promise<string | null> {
    try {
        const tabs = await chrome.tabs.query({ active: true, currentWindow: true })
        const tab = tabs[0]

        if (tab.url === undefined) {
            return null
        }

        if (!isStreamingVideo(tab.url)) {
            return null
        }

        return tab.url
    } catch (error) {
        console.error(error)
        return null
    }
}

function isStreamingVideo (url: string): boolean {
    return STREAMING_VIDEO_PLATFORMS.some(substring => url.includes(substring))
}
