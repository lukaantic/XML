import React, {useEffect} from 'react';

import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';

import { useNavigate } from 'react-router-dom';

import "./Header.scss"

export default function Header ({onSliderOpen, isOpen}) {
    const Navigate = useNavigate();

    const handleLogoutClick = () => {
        localStorage.clear();
        Navigate('/'); 
    }

    const handleRegisterClick = () => { 
        Navigate('/register')
    }

    const handleProfileClick = async () => {
    }

    return (
        <div className='header-container'>
            <Box id='header' sx={{ flexGrow: 1 }}>
                <AppBar position="static">
                    <Toolbar>
                        <IconButton
                            size="large"
                            edge="start"
                            color="inherit"
                            aria-label="menu"
                            sx={{ mr: 2 }}
                            onClick={onSliderOpen}
                        >
                            <MenuIcon />
                        </IconButton>
                            <Typography id='logo' variant="h6" component="div" sx={{ flexGrow: 1 }} className={isOpen && "move-logo"}>
                                Dislinkt
                            </Typography>
                            
                        <Button onClick={handleProfileClick} color="inherit">Profile</Button>
                        <Button onClick={handleLogoutClick} color="inherit">Log out</Button>
                        <Button onClick={handleRegisterClick} color="inherit">Register out</Button>
                    </Toolbar>
                </AppBar>
            </Box>
        </div>
    )
}

