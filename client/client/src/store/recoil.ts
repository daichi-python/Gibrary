import { atom } from "recoil";

export const AuthFormState = atom({
    key: "authFormState",
    default: {
        isSignupForm: true,
        sendMail: false,
    }
})

export const UserInformation = atom({
    key: "userInformation",
    default: {
        Name: "",
        Image: ""
    }
})