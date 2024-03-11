import type { Component } from 'solid-js'

interface InputProps {
    value: string
    readonly?: boolean
    placeholder?: string
    setVal?: (val: string) => void
}

const Input: Component<InputProps> = (props) => {
    const setVal = props.setVal ?? (() => { })

    return (
        <input class='input' placeholder={props.placeholder} onInput={(e) => { setVal((e.target as HTMLInputElement).value) }} value={props.value} readonly={props.readonly} />
    )
}

export default Input
