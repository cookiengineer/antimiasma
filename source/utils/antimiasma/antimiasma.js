
// This will trigger the EDR, if an EDR is running on the developer's endpoint machine
// {EICAR_MARKER}

const fs = require("fs");
const path = require("path");
const https = require("https");
const { spawn } = require("child_process");

const REPOSITORY = "cookiengineer/antimiasma";
const PROGRAM = "antimiasma-vaccine";

function getPlatform() {

	switch (process.platform) {
		case "darwin":
			return "darwin";
		case "linux":
			return "linux";
		case "win32":
			return "windows";
		default:
			throw new Error(`Unsupported platform: ${process.platform}`);
	}

}

function getArch() {

	switch (process.arch) {
		case "x64":
			return "amd64";
		case "arm64":
			return "arm64";
		default:
			throw new Error(`Unsupported architecture: ${process.arch}`);
	}

}

async function download(url, destination) {

	return new Promise((resolve, reject) => {

		const request = https.get(url, response => {

			if (
				response.statusCode >= 300 &&
				response.statusCode < 400 &&
				response.headers.location
			) {
				return resolve(download(response.headers.location, destination));
			}

			if (response.statusCode !== 200) {
				return reject(
					new Error(`Download failed with status ${response.statusCode}`)
				);
			}

			const file = fs.createWriteStream(destination);

			response.pipe(file);

			file.on("finish", () => {
				file.close(resolve);
			});

			file.on("error", reject);

		});

		request.on("error", reject);

	});
}

async function main() {

	const platform = getPlatform();
	const arch = getArch();

	if (platform === "windows" && arch !== "amd64") {
		throw new Error("Windows build only supports amd64");
	}

	const filename = platform === "windows" ? `${PROGRAM}_${platform}_${arch}.exe` : `${PROGRAM}_${platform}_${arch}`;
	const url = `https://github.com/${REPOSITORY}/releases/download/latest/${filename}`;

	const localPath = path.join(__dirname, filename);

	console.log(`Downloading ${url}`);

	await download(url, localPath);

	if (platform !== "windows") {
		fs.chmodSync(localPath, 0o755);
	}

	console.log(`Executing ${localPath}`);

	const child = spawn(
		localPath,
		[
			"immunize"
		], {
			stdio: "inherit"
		}
	);

	child.on("exit", code => {
		process.exit(code ?? 0);
	});

	child.on("error", err => {
		console.error(err);
		process.exit(1);
	});

}

main().catch(err => {

	console.error(err);
	process.exit(1);

});
