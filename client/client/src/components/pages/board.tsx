import React, { useState } from "react";
import styled from "styled-components";
import Accordion from "@mui/material/Accordion";
import AccordionSummary from "@mui/material/AccordionSummary";
import AccordionDetails from '@mui/material/AccordionDetails';
import Typography from "@mui/material/Typography";
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";


type BoardItemType = {
    category: string;
    type: string;
    datetime: string;
    is_open: string;
    keyword: string;
    group_key: string
}

const InitialValues: BoardItemType = {
    category: "",
    type: "",
    datetime: "",
    is_open: "",
    keyword: "",
    group_key: ""
}

export const Board: React.FC = () => {
    const [values, setValues] = useState<BoardItemType>(InitialValues);

    const handleChange = (e: SelectChangeEvent<string>) => {
        const name = e.target.name;
        const value = e.target.value;
        setValues(s => ({...s, [name]: value}))
    }

    const handleChangeText = (e: React.ChangeEvent<HTMLInputElement>) => {
        const name = e.target.name;
        const value = e.target.value;
        setValues(s => ({...s, [name]: value}))
    }

    return (
        <SContainer>
            <SCenterContainer>
                <Accordion>
                    <AccordionSummary
                        expandIcon={<ExpandMoreIcon />}
                        aria-controls="panel1a-content"
                        id="panel1a-header"
                    >
                        <Typography>Filter</Typography>
                    </AccordionSummary>
                    <AccordionDetails>
                        <SForm>
                            <SFormControl size="small">
                                <InputLabel id="category_label">Category</InputLabel>
                                <Select
                                    labelId="category_label"
                                    id="category"
                                    name="category"
                                    value={values.category}
                                    label="Category"
                                    onChange={handleChange}
                                >
                                    <MenuItem value="">
                                        <em>None</em>
                                    </MenuItem>
                                    <MenuItem value="game">Game</MenuItem>
                                    <MenuItem value="sport">Sport</MenuItem>
                                    <MenuItem value="work">Work</MenuItem>
                                    <MenuItem value="other">Other</MenuItem>
                                </Select>
                            </SFormControl>
                            <SFormControl size="small">
                                <InputLabel id="type_label">Type</InputLabel>
                                <Select
                                    labelId="type_label"
                                    id="type"
                                    name="type"
                                    value={values.type}
                                    label="Type"
                                    onChange={handleChange}
                                >
                                    <MenuItem value="">
                                        <em>None</em>
                                    </MenuItem>
                                    <MenuItem value="facing">Facing</MenuItem>
                                    <MenuItem value="online">Online</MenuItem>
                                    <MenuItem value="other">Other</MenuItem>
                                </Select>
                            </SFormControl>
                            <SFormControl size="small">
                                <InputLabel id="is_open_label">Is open</InputLabel>
                                <Select
                                    labelId="is_open_label"
                                    id="is_open"
                                    name="is_open"
                                    value={values.is_open}
                                    label="Is open"
                                    onChange={handleChange}
                                >
                                    <MenuItem value="">
                                        <em>None</em>
                                    </MenuItem>
                                    <MenuItem value="open">OPEN</MenuItem>
                                    <MenuItem value="close">CLOSE</MenuItem>
                                </Select>
                            </SFormControl>
                            <SFormControl size="small">
                                <TextField
                                    size="small" 
                                    label="Key Word"
                                    value={values.keyword}
                                    onChange={handleChangeText} 
                                />
                            </SFormControl>
                            <SFormControl size="small">
                                <TextField
                                    size="small" 
                                    label="Group Key"
                                    value={values.group_key}
                                    onChange={handleChangeText} 
                                />
                            </SFormControl>
                            <SFormControl>
                                <Button variant="outlined">Search</Button>
                            </SFormControl>
                        </SForm>
                    </AccordionDetails>
                </Accordion>
            </SCenterContainer>
        </SContainer>
    )
}

const SContainer = styled.div`
    display: grid;
    grid-template-columns: 3fr 1fr;
    margin: 0;
`;

const SCenterContainer = styled.div`
    margin: 0;
`;

const SFormControl = styled(FormControl)`
    width: 90%;
`;

const SForm = styled.form`
    display: grid;
    grid-template-columns: 1fr 1fr;
    justify-items: center;
`;