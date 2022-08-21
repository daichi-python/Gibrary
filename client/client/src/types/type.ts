import React from "react"

export type AuthFormStateType = {
    name: string
    email: string
    password: string
}

export type TextInputState = {
    name: string
    value: string
    type: string
    onChange: (e: React.ChangeEvent<HTMLInputElement>) => void
}

