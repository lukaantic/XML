import React, { useState, useContext } from 'react';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import Header from '../../../components/navigation/Header';
import { AuthenticationContext } from '../../../context/AuthenticationContext';
import { useNavigate } from 'react-router-dom';


const theme = createTheme();

const Registration = () => {
  const Navigate = useNavigate();

  localStorage.removeItem('token');

  const { signUp } = useContext(AuthenticationContext);

  const [signUpData, setSignUpData] = useState({
    email: '',
    password: '',
    name: '',
    surname: '',
    confirmPassword: '',
    username: '',
    biography: '',
    phoneNumber: '',
  });
  const [errorMessage, setErrorMessage] = useState();

  const isButtonDisabled =
    signUpData.email === '' ||
    signUpData.surname === '' ||
    signUpData.name === '' ||
    signUpData.username === '' ||
    signUpData.biography === '' ||
    signUpData.phoneNumber === '' ||
    signUpData.password === '' ||
    signUpData.confirmPassword === '' ||
    signUpData.password.length < 8 ||
    signUpData.password !== signUpData.confirmPassword;

  const successCallback = (response) => {

    if (response === "Created") {
      Navigate('/')
    }
  };

  const errorCallback = error => {
    alert(error)
  };

  const handleChange = (name, value) => {
    setSignUpData({ ...signUpData, [name]: value });
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    signUp(signUpData, successCallback, errorCallback);
  };


  return (
    <div>
      <Header />
      <ThemeProvider theme={theme}>
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
            <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
              <LockOutlinedIcon />
            </Avatar>
            <Typography component="h1" variant="h5">
              Sign up
            </Typography>
            <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 3 }}>
              <Grid container spacing={2}>
                <Grid item xs={12} sm={6}>
                  <TextField
                    name="name"
                    required
                    fullWidth
                    id="name"
                    label="First Name"
                    autoFocus
                    onChange={(event) => handleChange('name', event.target.value)}
                  />
                </Grid>
                <Grid item xs={12} sm={6}>
                  <TextField
                    required
                    fullWidth
                    id="surname"
                    label="Last Name"
                    name="surname"
                    onChange={(event) => handleChange('surname', event.target.value)}
                  />
                </Grid>
                <Grid item xs={12} >
                  <TextField
                    required
                    fullWidth
                    id="username"
                    label="Username"
                    name="username"
                    onChange={(event) => handleChange('username', event.target.value)}
                  />
                </Grid>
                <Grid item xs={12}>
                  <TextField
                    required
                    fullWidth
                    id="email"
                    label="Email Address"
                    name="email"
                    onChange={(event) => handleChange('email', event.target.value)}
                  />
                </Grid>
                <Grid item xs={12} >
                  <TextField
                    required
                    fullWidth
                    id="phoneNumber"
                    label="Phone number"
                    name="phoneNumber"
                    onChange={(event) => handleChange('phoneNumber', event.target.value)}
                  />
                </Grid>
                <Grid item xs={12} >
                  <TextField
                    required
                    fullWidth
                    id="biography"
                    label="Biography"
                    name="biography"
                    onChange={(event) => handleChange('biography', event.target.value)}
                  />
                </Grid>
                <Grid item xs={12}>
                  <TextField
                    required
                    fullWidth
                    name="password"
                    label="Password"
                    type="password"
                    id="password"
                    onChange={(event) => handleChange('password', event.target.value)}
                  />
                </Grid>
                <Grid item xs={12}>
                  <TextField
                    required
                    fullWidth
                    name="confirmPassword"
                    label="Confirm Password"
                    type="password"
                    id="confirmPassword"
                    onChange={(event) => handleChange('confirmPassword', event.target.value)}
                  />
                </Grid>
              </Grid>
              <Button
                disabled={isButtonDisabled}
                type="submit"
                fullWidth
                variant="contained"
                sx={{ mt: 3, mb: 2 }}
              >
                Sign Up
              </Button>
              <Grid container justifyContent="flex-end">
                <Grid item>
                  <Link href="/" variant="body2">
                    Already have an account? Sign in
                  </Link>
                </Grid>
              </Grid>
            </Box>
          </Box>
        </Container>
      </ThemeProvider>
    </div>
  );
}

export default Registration;