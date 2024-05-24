import Title from '../Title/Title';
import TextGradient from '../TextGradient/TextGradient';
import SideMenuItem from './SideMenuItem';

function SideMenu() {
  return (
    <div
      style={{
        display: 'flex',
        flex: 1,
        maxWidth: 220,
        backgroundColor: '#f1f1f1',
        borderRight: '1px solid #e1e1e1',
      }}
    >
      <div style={{ padding: 20, paddingTop: 35, flex: 1 }}>
        <Title style={{ marginBottom: 50 }}>
          <TextGradient>Sierra</TextGradient>
        </Title>
        <SideMenuItem route="/" name="Explore" />
        <SideMenuItem route="/sources" name="Sources" />
      </div>
    </div>
  );
}

export default SideMenu;
