import React, { useState, useContext, useEffect } from 'react';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { useNavigate } from 'react-router-dom';
import Header from '../../components/navigation/Header';
import { axiosInstance } from '../../axios/axios';
import { Checkbox } from '@mui/material';


const theme = createTheme();

const UserProfile = () => {
  const usersUsername = localStorage.getItem('username');

  const [userData, setUSetData] = useState();
  const [followers, setFollowers] = useState([]);
  const [checked, setChecked] = useState(false);

  const handleCheck = (event) => {
    setChecked(event.target.checked);
  }

  const getUser = async () => {
    try {
      const response = await axiosInstance.get(`/user/find-user/${usersUsername}`);

      setUSetData(response?.data);

      let privacy = response?.data?.profilePrivacy?.PrivacyType == 0 ? false : true;

      setChecked(privacy);

    } catch (error) {
      alert("Error: " + error.message);
    }
  }

  const getFollowers = async () => {
    try {
      const response = await axiosInstance.get(`/follow/followers/${userData?._id}`);

      setFollowers(response?.data);

    } catch (error) {
      alert("Error: " + error.message);
    }
  }

  useEffect(() => {
    getUser();
  }, []) 

  useEffect(() => {
    if (followers.length == 0) {
      getFollowers();
    }
  }, [followers]) 

  const [usersData, setUsersData] = useState({
    name: '',
    surname: '',
    skills: '',
    username: '',
    interests: '',
    experience: '',
  });

  const handleChange = (name, value) => {
    setUsersData({ ...usersData, [name]: value });
  };

  const updateUser = async () => {
    try {

      const payload = {
        _id: userData?._id,
        name: usersData?.name == "" ? userData?.name : usersData?.name,
        surname: usersData?.surname == "" ? userData?.surname : usersData?.surname,
        skills: usersData?.skills == "" ? userData?.skills : usersData?.skills,
        username: usersData?.username == "" ? userData?.username : usersData?.username,
        interests: usersData?.interests == "" ? userData?.interests : usersData?.interests,
        experience: usersData?.experience == "" ? userData?.experience : usersData?.experience,
      }

      const response = await axiosInstance.put("/user/update", payload);

    } catch (error) {
      alert("Error: " + error.message);
    }
  }

  const updatePrivacy = async () => {

    let privacy = checked ? 1 : 0;

    const payload = {
      _id: userData?._id,
      privacytype: privacy
    }
    const response = await axiosInstance.put("/user/update-profile-privacy", payload);
  }

  const handleSubmit = () => {
    updatePrivacy();
    updateUser();
    window.location.reload();
  };


  return (
    <div>
      <Header />
      <ThemeProvider theme={theme}>
        <Container component="main" maxWidth="xs">
          <Box
            sx={{
              marginTop: 8,
              display: 'flex',
              flexDirection: 'column',
              alignItems: 'center',
            }}
          >
            <Typography component="h1" variant="h5" style={{marginBottom:22}}>
              Your Profile
            </Typography>

            <Typography component="h1" variant="h5">
              Followers
            </Typography>
            {   followers?.length > 0 &&
                followers?.map((follower, index) => {
                  return (
                    <div key={index} >
                      {follower?.username}
                    </div>
                  )
                })
              }
            <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 3 }}>
              <Grid container spacing={2}>
                <Grid item xs={12} sm={6}>
                  <TextField
                    name="name"
                    required
                    fullWidth
                    id="name"
                    label={userData?.name}
                    autoFocus
                    placeholder={"First Name"}
                    onChange={(event) => handleChange('name', event.target.value)}
                  />
                </Grid>
                <Grid item xs={12} sm={6}>
                  <TextField
                    required
                    fullWidth
                    id="surname"
                    label={userData?.surname}
                    placeholder='Last Name'
                    name="surname"
                    onChange={(event) => handleChange('surname', event.target.value)}
                  />
                </Grid>
                <Grid item xs={12} >
                  <TextField
                    required
                    fullWidth
                    id="username"
                    label={userData?.username}
                    placeholder='Username'
                    name="username"
                    onChange={(event) => handleChange('username', event.target.value)}
                  />
                </Grid>
                <Grid item xs={12}>
                  <TextField
                    required
                    fullWidth
                    id="skills"
                    label={userData?.skills}
                    placeholder='Skills'
                    name="skills"
                    onChange={(event) => handleChange('skills', event.target.value)}
                  />
                </Grid>
                <Grid item xs={12} >
                  <TextField
                    required
                    fullWidth
                    id="interests"
                    label={userData?.interests}
                    placeholder='Interests'
                    name="interests"
                    onChange={(event) => handleChange('interests', event.target.value)}
                  />
                </Grid>
                <Grid item xs={12} >
                  <TextField
                    required
                    fullWidth
                    id="experience"
                    label={userData?.experience}
                    placeholder='Experience'
                    name="experience"
                    onChange={(event) => handleChange('experience', event.target.value)}
                  />
                </Grid>
                <Grid item xs={12} >
                  <TextField
                    required
                    fullWidth
                    id="education"
                    label={userData?.education}
                    placeholder='Education'
                    name="education"
                    onChange={(event) => handleChange('education', event.target.value)}
                  />
                </Grid>
              </Grid>
              <div style={{ marginTop:20}}>
                Is Private
              </div>
              <Checkbox
                checked={checked}
                onChange={handleCheck}
                inputProps={{ 'aria-label': 'controlled' }}
                style={{backgroundColor :'blue', marginTop:20}}
              />
              <Button
                type="submit"
                fullWidth
                variant="contained"
                sx={{ mt: 3, mb: 2 }}
              >
                Update
              </Button>

              
            </Box>
          </Box>
        </Container>
      </ThemeProvider>
    </div>
  );
}

export default UserProfile;