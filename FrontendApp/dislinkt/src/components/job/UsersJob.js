import * as React from 'react';
import Card from '@mui/material/Card';
import CardHeader from '@mui/material/CardHeader';
import CardMedia from '@mui/material/CardMedia';
import CardContent from '@mui/material/CardContent';
import CardActions from '@mui/material/CardActions';
import Avatar from '@mui/material/Avatar';
import IconButton from '@mui/material/IconButton';
import { red } from '@mui/material/colors';
import FavoriteIcon from '@mui/icons-material/Favorite';
import { Link } from '@mui/material';


const UsersJob = ({ title, urls, image, userName }) => {
  return (
    <Card sx={{ maxWidth: 345 }} style={{ backgroundColor: "#99CCEE", width: 400, marginBottom: 30 }}>
      <CardHeader
        avatar={
          <Avatar sx={{ bgcolor: red[500] }} aria-label="recipe" />
        }
        title={title}
        subheader={userName}
      />
      <CardMedia
        component="img"
        height="194"
        image={image}
      />
      <CardContent>
        {
          urls?.map(({ url }, index) => {
            <Link key={index} href={url}>{url}</Link>
          })
        }
      </CardContent>
      <CardActions disableSpacing>
        <IconButton aria-label="add to favorites">
          <FavoriteIcon />
        </IconButton>
      </CardActions>
    </Card>
  );
}

export default UsersJob;