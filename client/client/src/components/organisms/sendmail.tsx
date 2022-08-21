import React from "react";
import styled from "styled-components";
import MarkEmailReadIcon from '@mui/icons-material/MarkEmailRead';

export const SendMail: React.FC = () => {
    return (
        <SContainer>
            <STitleBox>
                <STitle>Email Verification</STitle>
            </STitleBox>
            <SIconBox>
                <MarkEmailReadIcon color="info" fontSize="large" />
            </SIconBox>
            <SMessageBox>
                <SMessage>Email has been sent to your registered email address.<br /><br />Please click on the link in the email to complete your registration.</SMessage>
            </SMessageBox>
        </SContainer>
    )
}

const SContainer = styled.div`
    width: 50%;
    background-color: white;
    border: 1px solid #6b6b6b3d;
    padding: 1rem;
`;

const STitleBox = styled.div`
    width: 100%;
    padding-left: 1rem;
`;

const STitle = styled.h2`
    font-family: Georgia;
    font-size: 1.8rem;
`;

const SIconBox = styled.div`
    display: flex;
    justify-content: center;
    align-items: center;
`

const SMessageBox = styled.div`
    display: flex;
    justify-content: center;
`;

const SMessage = styled.p`
    font-family: Georgia;
    font-size: 1.4rem;
    width: 80%;
`