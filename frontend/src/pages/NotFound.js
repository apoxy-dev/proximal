import { MinusCirlce } from 'iconsax-react';
import { Container, Grid, Paper, Typography } from '@mui/material';

const NotFound = () => (
  <Container>
    <Paper
      sx={{
        my: 1,
        mx: 'auto',
        p: 2,
      }}
    >
      <Grid container wrap="nowrap" spacing={2}>
        <Grid item>
          <MinusCirlce
            size="32"
            color="#FF8A65"
          />
        </Grid>
        <Grid item xs>
          <Typography
            variant="h5"
          >404</Typography>
          <Typography>Not Found</Typography>
        </Grid>
      </Grid>
    </Paper>
  </Container>
);

export default NotFound;
