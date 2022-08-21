import React, {useState} from "react";
import styled from "styled-components";
import { useNavigate } from "react-router-dom";
import { useRecoilState } from "recoil";
import { useTranslation } from "react-i18next";
import { 
    createUserWithEmailAndPassword, 
    GoogleAuthProvider,
    FacebookAuthProvider,
    sendEmailVerification, 
    signInWithPopup,
    setPersistence,
    inMemoryPersistence
} from "firebase/auth";
import { Button } from "@mui/material";
import { CircularProgress } from "@mui/material";

import { auth } from "../../auth/firebase";
import { AuthTextInput } from "../atoms/AuthTextInput";
import { AuthFormStateType } from "../../types/type";
import { AuthFormState } from "../../store/recoil";
import { SocialButton } from "../atoms/SocialButton";
import { PostUserInformation } from "../../api/init";


const initialFormState = {
    name: "",
    email: "",
    password: "",
}

export const Signup: React.FC = () => {
    // Value for sign up
    const [values, setValues] = useState<AuthFormStateType>(initialFormState);
    // Value to switch screens from recoil.
    const [authFormState, setAuthFormState] = useRecoilState(AuthFormState);
    // Value to toggle button display.
    const [isInProcess, setIsInProcess] = useState(false);

    const navigate = useNavigate();
    const { t } = useTranslation();

    const switchForm = () => {
        setAuthFormState({isSignupForm: !authFormState.isSignupForm, sendMail: false});
    }

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const name = e.target.name;
        const value = e.target.value;
        setValues(s => ({...s, [name]: value}))
    }

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        setIsInProcess(true);

        const { name, email, password } = values;
        console.log(`name: ${name}\nemail: ${email}\npassword: ${password}`)

        const actionCodeSettings = {
            url: 'http://localhost:3000',
            handleCodeInApp: true,
        };

        try {
            createUserWithEmailAndPassword(auth, email, password)
            .then((res) => {
                const user = res.user;
                sendEmailVerification(user, actionCodeSettings)
                .then(() => {
                    setAuthFormState({isSignupForm: true, sendMail: true});
                    PostUserInformation(name, email)
                }).catch((e) => {
                    alert("Failed to send email.")
                    console.log(e)
                })
            }).catch((e) => {
                alert("")
                setIsInProcess(false);
                console.log(e);
            })
        } catch (e) {
            setIsInProcess(false);
            alert("Failed to create user.")
            console.log(e)
        }
    }

    const GoogleSignup = (e: React.MouseEvent<HTMLInputElement>) => {
        const provider = new GoogleAuthProvider();

        setPersistence(auth, inMemoryPersistence)
        .then(() => {
            return (
                signInWithPopup(auth, provider)
                .then((res) => {
                    const credential = GoogleAuthProvider.credentialFromResult(res);
                    const token = credential?.accessToken;
                    const user = res.user;
                    
                    if (user.displayName && user.email) {
                        PostUserInformation(user.displayName, user.email);
                        navigate("/");
                    }
                })
                .catch((e) => {
                    alert("Failed to log in with google");
                    console.log(e);
                })
            )
        })
    }

    const FacebookSignup = (e: React.MouseEvent<HTMLInputElement>) => {
        const provider = new FacebookAuthProvider();

        setPersistence(auth, inMemoryPersistence)
        .then(() => {
            return (
                signInWithPopup(auth, provider)
                .then((res) => {
                    const credential = FacebookAuthProvider.credentialFromResult(res);
                    const token = credential?.accessToken;
                    const user = res.user;
                    
                    if (user.displayName && user.email) {
                        PostUserInformation(user.displayName, user.email);
                        navigate("/");
                    }
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
                <STitle>{t("signup.title")}</STitle>
            </STitleBox>
            <div>
                <SSocialButtonBox>
                    <SocialButton onClick={GoogleSignup} provider="google" message={t("signup.google")}/>
                </SSocialButtonBox>
                <SSocialButtonBox>
                    <SocialButton onClick={FacebookSignup} provider="facebook" message={t("signup.facebook")} />
                </SSocialButtonBox>
            </div>
            <SSeparate>or</SSeparate>
            <SForm action="POST" onSubmit={handleSubmit}>
                <SInputBox>
                    <AuthTextInput 
                        name="name"
                        value={values.name}
                        type="text"
                        onChange={handleChange}
                    />
                </SInputBox>
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
                    {isInProcess || <Button type="submit" variant="contained">{t("signup.button")}</Button>}
                </SButtonBox>
            </SForm>
            <SSwitchBox>
                <p>{t("signup.switch")} <Button size="small" onClick={switchForm}>{t("signup.switchButton")}</Button></p>
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

const SForm = styled.form`
    margin: 0;
`;

const STitleBox = styled.div`
    width: 50%;
    padding-left: 1rem;
`;

const STitle = styled.h2`
    font-family: Georgia, "游ゴシック";
    font-size: 1.8rem;
`;

const SSocialButtonBox = styled.div`
    margin: 0;
    margin-bottom: 0.3rem;
    display: flex;
    justify-content: center;
`;

const SInputBox = styled.div`
    margin: 0;
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
    align-items: center;
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
        margin: 0 1.5em;
    }
`;