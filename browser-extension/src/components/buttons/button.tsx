import type { Component } from 'solid-js'

interface ButtonProps {
    onClick?: () => void
    type?: 'button' | 'submit'
    value: string
}

const Button: Component<ButtonProps> = (props) => {
    return (
        <button type={props.type} onClick={props.onClick} class='button button-primary'>
            {props.value}
        </button>
    )
}

export default Button
