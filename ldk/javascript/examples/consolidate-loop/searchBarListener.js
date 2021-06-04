import { ui, whisper, filesystem, network } from '@oliveai/ldk';
import { Patient, patientInfoFileName } from './Patient';

const { Message, Link, Markdown } = whisper.WhisperComponentType;

const emitResultWhisper = async (rows) => {
  const searchResults = [];
  rows.forEach((row) => {
    searchResults.push({
      text: `${row.firstName} ${row.lastName} ${row.email}`,
      onClick: () => {
        console.log(row.serialize());

        whisper.create({
          label: 'Consolidate Loop - Patient Search Result',
          onClose: () => {
            console.log('Patient Search Result is closed');
          },
          components: [
            {
              header: `${row.firstName} ${row.lastName}`,
              body: row.patientInfo.reduce(
                (rst, k) => `
${rst}
${k}:   ${row[k]}
                            `,
                '',
              ),
              type: Message,
            },
          ],
        });
      },
      type: Link,
    });
  });

  await whisper.create({
    label: 'Consolidate Loop - Patient Search Result',
    onClose: () => {
      console.log('Patient Search Result');
    },
    components:
      searchResults.length > 0
        ? searchResults
        : [
            {
              body: `No results are found matching provided criteria`,
              type: Markdown,
            },
          ],
  });
};

const getPatients = async () => {
  const patients = [];
  try {
    if (!filesystem.exists(patientInfoFileName)) {
      console.warn('No patients file exists');
      return patients;
    }

    const data = await filesystem.readFile(patientInfoFileName);
    const decodedData = await network.decode(data);
    decodedData.split('|').forEach((row) => {
      if (row) {
        const splitRow = row.split(':');
        const newPatient = new Patient({
          firstName: splitRow[0],
          lastName: splitRow[1],
          dob: splitRow[2],
          gender: splitRow[3],
          telephone: splitRow[4],
          email: splitRow[5],
          visitReason: splitRow[6],
          appointmentDate: splitRow[7],
          appointmentTime: splitRow[8],
        });
        patients.push(newPatient);
      }
    });

    return patients;
  } catch (error) {
    console.error(error);
  }

  return patients;
};

export const start = async () => {
  await ui.listenSearchbar(async (value) => {
    const rows = [];
    const patients = await getPatients();
    patients.forEach((p) => {
      if (p.getSearchCred().includes(value)) {
        rows.push(p);
      }
    });

    await emitResultWhisper(rows);
  });
};
