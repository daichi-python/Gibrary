import React from "react";
import styled from "styled-components";
import { Checkbox } from "@mui/material";
import FavoriteBoader from "@mui/icons-material/FavoriteBorder";
import Favorite from "@mui/icons-material/Favorite";
import ReplyIcon from "@mui/icons-material/Reply";
import ShareIcon from '@mui/icons-material/Share';


const label = { inputProps: { "aria-label": "Checkbox demo" } };

export const HomeActions: React.FC = () => {
    return (
        <SContainer>
            <SButtonBox>
                <Checkbox {...label} icon={<FavoriteBoader />} checkedIcon={<Favorite color="error" />} />
            </SButtonBox>
            <SButtonBox>
                <SReplyIcon />
            </SButtonBox>
            <SButtonBox>
                <SShareIcon />
            </SButtonBox>
        </SContainer>
    )
}

const SContainer = styled.div`
    display: flex;
    align-items: center;
    justify-content: space-between;
`;

const SButtonBox = styled.div`
    display: flex;
    align-items: center;
    margin: 0 3rem;
`;

const SReplyIcon = styled(ReplyIcon)`
    cursor: pointer;
`;

const SShareIcon = styled(ShareIcon)`
    cursor: pointer;
`