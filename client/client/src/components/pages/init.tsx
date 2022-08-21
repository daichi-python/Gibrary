import React from "react";
import styled from "styled-components";
import { useRecoilValue } from "recoil";

import { AuthFormState } from "../../store/recoil";
import { SendMail } from "../organisms/sendmail";

type InitialForms = {
    Signup: React.FC;
    Signin: React.FC;
}

export const Init: React.FC<InitialForms> = (props) => {
    const {Signup, Signin} = props;
    const authFormState = useRecoilValue(AuthFormState);

    return (
        <SContainer>
            <SInnerContainer>
                {authFormState.isSignupForm && (authFormState.sendMail ? <SendMail /> : <Signup />)}
                {authFormState.isSignupForm || <Signin />}
            </SInnerContainer>
            <SInnerContainer>
                <h2>Auth Page</h2>
            </SInnerContainer>
        </SContainer>
    )
}

const SContainer = styled.div`
    margin: 1rem 2rem;
    background-color: red;
    display: grid;
    grid-template-columns: 1fr 1fr;
    height: 95vh;
`;

const SInnerContainer = styled.div`
    margin: 2rem 1rem;
    background-color: yellow;
    display: flex;
    justify-content: center;
    align-items: center;
`;