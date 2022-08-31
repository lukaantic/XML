import React, { useContext, useEffect } from 'react';

import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';

import { useNavigate } from 'react-router-dom';

import "./Header.scss"
import { AuthenticationContext } from '../../context';

export default function Header() {
  const Navigate = useNavigate();

  const { token, deleteUserToken } = useContext(AuthenticationContext);

  const handleLogoutClick = () => {
    deleteUserToken();
    Navigate('/');
  }

  const handleRegisterClick = () => {
    Navigate('/registration')
  }

  const handleProfileClick = () => {
  }

  const handleHomeClick = () => {
    Navigate('/home')
  }

  const handleLogInClick = () => {
    Navigate('/');
  }

  return (
    <div className='header-container'>
      <Box id='header' sx={{ flexGrow: 1 }}>
        <AppBar position="static">
          <Toolbar>
            <Typography id='logo' variant="h6" component="div" sx={{ flexGrow: 1 }}>
              Dislinkt
            </Typography>

            <Button onClick={handleHomeClick} color="inherit">Home</Button>
            {token &&
              <Button onClick={handleProfileClick} color="inherit">Profile</Button>
            }
            {token ?
              <Button onClick={handleLogoutClick} color="inherit">Log out</Button>
              :
              <Button onClick={handleLogInClick} color="inherit">Log In</Button>
            }
            {!token &&
              <Button onClick={handleRegisterClick} color="inherit">Register</Button>
            }
          </Toolbar>
        </AppBar>
      </Box>
    </div>
  )
}

