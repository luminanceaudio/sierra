import Title from '../Title/Title';
import TextGradient from '../TextGradient/TextGradient';

function SideMenu() {
  return (
    <div
      style={{
        display: 'flex',
        flex: 1,
        maxWidth: 220,
        backgroundColor: '#f1f1f1',
        // borderRadius: 22,
      }}
    >
      <div style={{ padding: 20, paddingTop: 35, flex: 1 }}>
        <Title>
          <TextGradient>Sierra</TextGradient>
        </Title>
      </div>
    </div>
  );
}

export default SideMenu;
