import { Link, useLocation } from 'react-router-dom';
import Title from '../Title/Title';

type SideMenuItemProps = {
  name: string;
  route: string;
};

function SideMenuItem({ name, route }: SideMenuItemProps) {
  const location = useLocation();

  return (
    <Link
      to={route}
      style={{
        display: 'block',
        margin: 0,
      }}
    >
      <Title
        style={{
          fontSize: 18,
          cursor: 'pointer',
          padding: '15px 0px',
          fontWeight: location.pathname === route ? 600 : 400,
        }}
      >
        {name}
      </Title>
    </Link>
  );
}

export default SideMenuItem;
