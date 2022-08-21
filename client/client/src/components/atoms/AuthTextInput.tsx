import React, { useState } from "react";
import IconButton from '@mui/material/IconButton';
import Input from '@mui/material/Input';
import InputLabel from '@mui/material/InputLabel';
import InputAdornment from '@mui/material/InputAdornment';
import FormControl from '@mui/material/FormControl';
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';

import { capitalize } from "../../utils/utils";
import { TextInputState } from "../../types/type";



export const AuthTextInput: React.FC<TextInputState> = (props) => {
    const {name, value, type, onChange} = props;
    const [showPassword, setShowPassword] = useState(false);

    const handleClickShowPassword = () => {
        setShowPassword(!showPassword);
    };
    
    const handleMouseDownPassword = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
    };

    const stringSize = (str: string) => {
        if (str.length === 0) return;
        return `${str.length.toString()}/30`;
    }

    return (
        <FormControl sx={{ m: 1, width: '25ch' }} variant="standard">
            <InputLabel htmlFor={`standard-adornment-${name}`}>{capitalize(name)}</InputLabel>
            <Input
                id={`standard-adornment-${name}`}
                name={name}
                type={name !== "password" ? type : (showPassword ? 'text' : 'password')}
                value={value}
                onChange={onChange}
                endAdornment={name !== "email" &&
                <InputAdornment position="end">
                    {name === "password" ? <IconButton
                        aria-label="toggle password visibility"
                        onClick={handleClickShowPassword}
                        onMouseDown={handleMouseDownPassword}
                    >
                    {showPassword ? <VisibilityOff /> : <Visibility />}
                    </IconButton> : stringSize(value)}
                </InputAdornment>
                }
            />
        </FormControl>
    )
}