import Layout from '../../components/Layout/Layout';
import Widget from '../../components/Widget/Widget';

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
        hello
      </Widget>
    </Layout>
  );
}

export default Main;
