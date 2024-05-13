import Layout from '../../components/Layout/Layout';
import Widget from '../../components/Widget/Widget';
import { useSources } from '../../../api/sources';
import { Source, SourceNew } from './Source';

function Sources() {
  const { data: sources } = useSources();

  return (
    <Layout title="Sources">
      <Widget
        style={{
          display: 'flex',
          flexDirection: 'row',
          paddingBottom: 100,
        }}
      >
        <div
          style={{
            display: 'flex',
            flex: 1,
            flexDirection: 'column',
            gap: 15,
            maxWidth: 600,
          }}
        >
          <SourceNew />
          {sources?.data?.sources?.map((source) => (
            <Source key={source.uri} uri={source.uri} />
          ))}
        </div>
      </Widget>
    </Layout>
  );
}

export default Sources;
