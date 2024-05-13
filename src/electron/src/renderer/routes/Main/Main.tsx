import Layout from '../../components/Layout/Layout';
import Widget from '../../components/Widget/Widget';
import Samples from '../../components/Samples/Samples';

function Main() {
  return (
    <Layout title="Samples">
      <Widget
        style={{
          display: 'flex',
          flex: 1,
          alignItems: 'center',
        }}
      >
        <Samples />
      </Widget>
    </Layout>
  );
}

export default Main;
