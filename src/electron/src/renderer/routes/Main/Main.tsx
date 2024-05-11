import Layout from '../../components/Layout/Layout';
import Widget from '../../components/Widget/Widget';
import Samples from '../../components/Samples/Samples';

function Main() {
  return (
    <Layout>
      <Widget
        title="List"
        style={{
          display: 'flex',
          flex: 1,
        }}
      >
        <Samples />
      </Widget>
    </Layout>
  );
}

export default Main;
