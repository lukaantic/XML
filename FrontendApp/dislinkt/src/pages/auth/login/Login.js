import React, { useState, useContext, useEffect } from 'react';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { useNavigate, Navigate } from 'react-router-dom';
import Header from '../../../components/navigation/Header';
// import { AuthenticationContext } from '../../../context/AuthenticationContext';

// import { userService } from '../../../services';

// import { GetTokenFromLocalStorage } from "../../../helper/Token";


const theme = createTheme();

export default function Login() {
  const navigate = useNavigate();
  //   const { login } = useContext(AuthenticationContext);

  const [loginData, setLoginData] = useState({
    email: '',
    password: '',
  });

  const handleChange = (name, value) => {
    setLoginData({ ...loginData, [name]: value });
  };

  //   useEffect(() => {
  //     let date = new Date()
  //     date.setHours(date.getHours() + 24)
  //     let loginTime = { value: "value", timestamp: date }
  //     localStorage.setItem("loginTime", JSON.stringify(loginTime));
  //   }, [])

  //   let token = GetTokenFromLocalStorage();

  //   if (token) {
  //     return <Navigate to="/home" replace />;
  //   }

  const loginUser = (event) => {
    event.preventDefault();

    // login(loginData, successCallback, errorCallback);
  };

  //   const successCallback = async () => {
  //     let result = await userService['getAdminProfile']();
  //     localStorage.setItem('loggedAdmin', JSON.stringify(result[0]));
  //     navigate('/home');
  //   };

  //   const errorCallback = (error) => {
  //     alert(error)
  //   };

  return (
    <ThemeProvider theme={theme}>
      <Header />
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}
        >
          <Typography component="h1" variant="h5">
            Sign in
          </Typography>
          <Box component="form" onSubmit={loginUser} noValidate sx={{ mt: 1 }}>
            <TextField
              margin="normal"
              required
              fullWidth
              id="email"
              label="Email Address"
              name="email"
              autoComplete="email"
              autoFocus
              onChange={(event) => handleChange('email', event.target.value)}
            />
            <TextField
              margin="normal"
              required
              fullWidth
              name="password"
              label="Password"
              type="password"
              id="password"
              autoComplete="current-password"
              onChange={(event) => handleChange('password', event.target.value)}
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              Sign In
            </Button>
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  );
}