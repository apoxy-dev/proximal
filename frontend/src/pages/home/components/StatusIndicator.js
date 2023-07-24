import { useTheme } from '@mui/material/styles';
import Typography from '@mui/material/Typography';

const StatusIndicator = ({ icon, children, variant }) => {
  const theme = useTheme();
  const color = variant === 'disabled' ? theme.palette.text.disabled : theme.palette.secondary.main;
  return (
    <Typography
      color="primary"
      sx={{
        display: 'flex',
        color: color,
        marginTop: '4px',
        fontWeight: 'bold',
        '& svg': {
          margin: '0 16px',
          color: color,
        }
      }}
    >
      {icon}
      {children}
    </Typography>
  );
};

export default StatusIndicator;
