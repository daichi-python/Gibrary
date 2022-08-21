import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import styled from "styled-components";
import { useRecoilState } from "recoil";
import { 
    signInWithEmailAndPassword, 
    signInWithPopup, 
    GoogleAuthProvider, 
    FacebookAuthProvider, 
    setPersistence,
    inMemoryPersistence
} from "firebase/auth";
import { Button } from "@mui/material";
import { CircularProgress } from "@mui/material";

import { AuthFormStateType } from "../../types/type";
import { AuthTextInput } from "../atoms/AuthTextInput";
import { AuthFormState } from "../../store/recoil";
import { SocialButton } from "../atoms/SocialButton";
import { auth } from "../../auth/firebase";


const initialFormState = {
    email: "",
    password: "",
}

export const Signin: React.FC = () => {
    // Value for sign in.
    const [values, setValues] = useState<Omit<AuthFormStateType, "name">>(initialFormState);
    // Value to switch screens from recoil.
    const [authFormState, setAuthFormState] = useRecoilState(AuthFormState);
    // Value to toggle button display.
    const [isInProcess, setIsInProcess] = useState(false);

    const navigate = useNavigate();

    const switchForm = () => {
        setAuthFormState({isSignupForm: !authFormState.isSignupForm, sendMail: false});
    }

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const name = e.target.name;
        const value = e.target.value;
        setValues(s => ({...s, [name]: value}))
    }

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setIsInProcess(true);
        const {email, password} = values;
        console.log(`email: ${email}\npassword: ${password}`);

        setPersistence(auth, inMemoryPersistence)
        .then(() => {
            return (
                signInWithEmailAndPassword(auth, email, password)
                .then((res) => {
                    const user = res.user;
                    if (user.emailVerified) {
                        navigate("/")
                    } else {
                        alert("Email verification has not yet been completed.")
                        setIsInProcess(false);
                    }
                })
                .catch((e) => {
                    alert("Failed to log in.\nPlease try again.")
                    console.log(e);
                    setIsInProcess(false);
                })
            )
        })
    }

    const GoogleSignin = (e: React.MouseEvent<HTMLInputElement>) => {
        const provider = new GoogleAuthProvider();

        setPersistence(auth, inMemoryPersistence)
        .then(() => {
            return (
                signInWithPopup(auth, provider)
                .then((res) => {
                    const credential = GoogleAuthProvider.credentialFromResult(res);
                    const token = credential?.accessToken;
                    const user = res.user;
                    navigate("/");
                })
                .catch((e) => {
                    alert("Failed to log in with google");
                    console.log(e);
                })
            )
        })
    }

    const FacebookSignin = (e: React.MouseEvent<HTMLInputElement>) => {
        const provider = new FacebookAuthProvider();

        setPersistence(auth, inMemoryPersistence)
        .then(() => {
            return (
                signInWithPopup(auth, provider)
                .then((res) => {
                    const credential = FacebookAuthProvider.credentialFromResult(res);
                    const token = credential?.accessToken;
                    const user = res.user;
                    navigate("/");
                })
                .catch((e) => {
                    alert("Failed to log in with facebook");
                    console.log(e);
                })
            )
        })
    }

    return (
        <SContainer>
            <STitleBox>
                <STitle>Sign in</STitle>
            </STitleBox>
            <div>
                <SSocialButtonBox>
                    <SocialButton onClick={GoogleSignin} provider="google" message="Sign in with Google." />
                </SSocialButtonBox>
                <SSocialButtonBox>
                    <SocialButton onClick={FacebookSignin} provider="facebook" message="Sign in with Facebook." />
                </SSocialButtonBox>
            </div>
            <SSeparate>or</SSeparate>
            <SForm action="POST" onSubmit={handleSubmit}>
                <SInputBox>
                    <AuthTextInput 
                        name="email"
                        value={values.email}
                        type="email"
                        onChange={handleChange}
                    />
                </SInputBox>
                <SInputBox>
                    <AuthTextInput
                        name="password"
                        value={values.password}
                        type="password"
                        onChange={handleChange}
                    />
                </SInputBox>
                <SButtonBox>
                    {isInProcess && <CircularProgress />}
                    {isInProcess || <Button type="submit" variant="contained">Sign in</Button>}
                </SButtonBox>
            </SForm>
            <SSwitchBox>
                <p>If you don't have account. <Button size="small" onClick={switchForm}>Sign up</Button></p>
            </SSwitchBox>
        </SContainer>
    )
}

const SContainer = styled.div`
    width: 50%;
    background-color: white;
    border: 1px solid #6b6b6b3d;
    padding: 1rem;
`;

const SSocialButtonBox = styled.div`
    margin: 0;
    margin-bottom: 0.3rem;
    display: flex;
    justify-content: center;
`;

const SForm = styled.form`
    margin: 0;
`;

const STitleBox = styled.div`
    width: 50%;
    padding-left: 1rem;
`;

const STitle = styled.h2`
    font-family: Georgia;
    font-size: 1.8rem;
`;

const SInputBox = styled.div`
    margin-top: 0.5;
    display: flex;
    justify-content: center;
`;

const SButtonBox = styled.div`
    display: flex;
    justify-content: center;
    margin-top: 0.8rem;
`;

const SSwitchBox = styled.div`
    display: flex;
    justify-content: right;
`;

const SSeparate = styled.p`
    font-size: 1.1rem;
    margin-top: 1rem;
    display: flex;
    align-items: center;
    justify-content: center;
    &:before, :after {
        border-top: 1px solid #8584847c;
        content: "";
        width: 9em;
        margin: 0 1em;
    }
`;