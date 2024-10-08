using System;
using System.Drawing;
using System.Windows.Forms;

namespace DisplayPngApp
{
    static class Program
    {
        [STAThread]
        static void Main()
        {
            Application.SetHighDpiMode(HighDpiMode.SystemAware);
            Application.EnableVisualStyles();
            Application.SetCompatibleTextRenderingDefault(false);
            Application.Run(new MainForm());
        }
    }

    public class MainForm : Form
    {
        private FlowLayoutPanel flowLayoutPanel;
        private string[] imagePaths =
        {
            "01-Apple.png",
            "02-Banana.png",
            "03-Blueberry.png",
            "04-Orange.png",
            "05-Strawberry.png",
            "06-Watermelon.png",
            "07-Seven.png",
        };

        public MainForm()
        {
            this.Text = "Display PNG Images";
            this.Width = 800;
            this.Height = 600;

            flowLayoutPanel = new FlowLayoutPanel
            {
                Dock = DockStyle.Fill,
                AutoScroll = true
            };

            Console.WriteLine("Reading files...");
            foreach (var imagePath in imagePaths)
            {
                var png = "C:/Users/joebo/Projects/2024/NETSlot/Reel-Images/" + imagePath;
                Console.WriteLine(png);
                PictureBox pictureBox = new PictureBox
                {
                    Image = Image.FromFile(png),
                    SizeMode = PictureBoxSizeMode.Zoom,
                    Width = 200,
                    Height = 200,
                    Margin = new Padding(10)
                };
                flowLayoutPanel.Controls.Add(pictureBox);
            }

            this.Controls.Add(flowLayoutPanel);
        }
    }
}