import React from "react";
import styled from "styled-components";

import { Avatar } from "@mui/material";

type Props = {
    Image: string
    Name: string
}

export const ImageAndName: React.FC<Props> = (props) => {
    const { Image, Name } = props;

    return (
        <SContainer>
            <SImageBox>
                <Avatar sx={{ width: 30, height: 30 }} src={Image} />
            </SImageBox>
            <SNameBox>
                <SName>{Name}</SName>
            </SNameBox>
        </SContainer>
    )
}

const SContainer = styled.div`
    display: flex;
    justify-content: left;
    align-items: center;
    margin-left: 1.3rem;
`;

const SImageBox = styled.div`
    margin-right: 0.8rem;
`;

const SNameBox = styled.div`
    margin: 0;
`;

const SName = styled.h4`
    font-family: Georgia, 'Times New Roman', Times, serif;
    font-size: 1.2rem;
`;
