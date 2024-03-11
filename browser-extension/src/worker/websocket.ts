export function newWebSocket (url: string, port: chrome.runtime.Port): WebSocket {
    const ws = new WebSocket(url)

    ws.onopen = () => {
        console.log('Kimparty: WebSocket connection opened')
    }

    ws.onclose = () => {
        console.log('Kimparty: WebSocket connection closed')
    }

    ws.onerror = (error) => {
        console.error('Kimparty: WebSocket error', error)
    }

    ws.onmessage = (message) => {
        const parsedMessage = JSON.parse(message.data as string)

        if (parsedMessage.content === undefined) {
            return
        }

        port.postMessage({ action: 'message', author: 'websocket', payload: { ...parsedMessage } })
    }

    return ws
}

export interface WebSocketMessage {
    id: string
    content: string
    member_id: string
    party_id: string
    type: number
    created_at: string
}
