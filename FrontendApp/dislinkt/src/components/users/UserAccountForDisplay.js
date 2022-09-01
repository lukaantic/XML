import  React, {useEffect, useState} from 'react';
import { styled } from '@mui/material/styles';
import Card from '@mui/material/Card';
import CardHeader from '@mui/material/CardHeader';
import CardMedia from '@mui/material/CardMedia';
import CardContent from '@mui/material/CardContent';
import CardActions from '@mui/material/CardActions';
import Avatar from '@mui/material/Avatar';
import { red } from '@mui/material/colors';
import { IconButton, Link, Typography } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import AccessibilityIcon from '@mui/icons-material/Accessibility';
import { axiosInstance } from '../../axios/axios';

const UserAccountForDisplay = ({ title, urls, image, username, name, surname, id }) => {
  const goToProfile = () => {
    window.location.href=`/account${username}`
  }

  const usersUsername = localStorage.getItem('username')

  const [userData, setUSetData] = useState();
  
  const getUser = async () => {
    try {
      const response = await axiosInstance.get(`/user/find-user/${usersUsername}`);
      setUSetData(response?.data);
    } catch (error) {
      alert("Error: " + error.message);
    }
  }

  useEffect(() => {
    getUser();
  }, []) 

  const handleFollowCLick = async () => {
    try {
      const payload = {
        followerId: id,
        followedId: userData?._id,
        isPrivate: false
      }

      const response = await axiosInstance.post("/follow/follow", payload);

    } catch (error) {
      alert("Error: " + error.message);
    }
  }


  return (
      <Card sx={{ maxWidth: 345 }} style={{ backgroundColor: "#99CCEE", width: 400, marginBottom: 30 }}>
        <div onClick={goToProfile}>
        <CardHeader
          avatar={
            <Avatar sx={{ bgcolor: red[500] }} aria-label="recipe" />
          }
          title={title}
          subheader={username}
        />
        <div style={{ display: 'flex', flexDirection: 'row' }}>
          <Typography style={{ marginLeft: 20, }}>
            {name}
          </Typography>
          <Typography style={{ marginLeft: 8, }}>
            {surname}
          </Typography>
        </div>
        <CardContent>
          {

          }
        </CardContent>
        </div>
        <CardActions disableSpacing>
          <IconButton aria-label="dislike" onClick={handleFollowCLick} >
            Follow <AccessibilityIcon />
          </IconButton>
        </CardActions>
      </Card>
  );
}

export default UserAccountForDisplay;