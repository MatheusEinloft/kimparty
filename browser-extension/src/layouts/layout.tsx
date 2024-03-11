import type { Component, JSX } from 'solid-js'
import styles from './layout.module.css'

const Layout: Component<{ title?: string, subtitle?: string, children?: JSX.Element }> = (props) => {
    return (
        <>
            <header>
                <h1 class={styles.title}>{props.title}</h1>
                <h3>{props.subtitle}</h3>
            </header>
            <main class={styles.section}>
                {props.children}
            </main>
        </>
    )
}

export default Layout
