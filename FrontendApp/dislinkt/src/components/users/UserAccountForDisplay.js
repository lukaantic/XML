import * as React from 'react';
import { styled } from '@mui/material/styles';
import Card from '@mui/material/Card';
import CardHeader from '@mui/material/CardHeader';
import CardMedia from '@mui/material/CardMedia';
import CardContent from '@mui/material/CardContent';
import CardActions from '@mui/material/CardActions';
import Avatar from '@mui/material/Avatar';
import { red } from '@mui/material/colors';
import { Link, Typography } from '@mui/material';
import { useNavigate } from 'react-router-dom';

const UserAccountForDisplay = ({ title, urls, image, username, name, surname, id }) => {
  const Navigate = useNavigate();

  const goToProfile = () => {
    Navigate(`/account${username}`)
  }

  return (
    <div onClick={goToProfile}>
      <Card sx={{ maxWidth: 345 }} style={{ backgroundColor: "#99CCEE", width: 400, marginBottom: 30 }}>
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
        <CardActions disableSpacing>
        </CardActions>
      </Card>
    </div>
  );
}

export default UserAccountForDisplay;