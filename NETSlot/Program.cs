using System;
using System.Drawing;
using System.IO;
using System.Linq;
using System.Windows.Forms;

public class MainForm : Form
{
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

    private System.Windows.Forms.Timer timer;
    private System.Windows.Forms.Timer durationTimer;
    private Random random;
    private Button restartButton;

    public MainForm()
    {
        this.Text = "Display PNG Images";
        this.Width = 800;
        this.Height = 600;
        this.FormBorderStyle = FormBorderStyle.FixedDialog; // Prevent resizing
        this.MaximizeBox = false; // Disable the maximize button

        random = new Random();

        timer = new System.Windows.Forms.Timer();
        timer.Interval = 100; // 100 ms
        timer.Tick += Timer_Tick;

        durationTimer = new System.Windows.Forms.Timer();
        durationTimer.Interval = 2000; // 2 seconds
        durationTimer.Tick += DurationTimer_Tick;

        restartButton = new Button
        {
            Text = "Spin",
            Location = new Point(400 - 50, 300),
            Width = 100,
            Height = 30
        };
        restartButton.Click += RestartButton_Click;

        this.Controls.Add(restartButton);

        // Start the timers initially
        StartTimers();
    }

    private void Timer_Tick(object? sender, EventArgs e)
    {
        var homePath = Environment.GetEnvironmentVariable("USERPROFILE");
        if (string.IsNullOrEmpty(homePath))
        {
            throw new InvalidOperationException("USERPROFILE environment variable is not set.");
        }

        Console.WriteLine("Reading files...");
        this.Controls.Clear(); // Clear previous images except the button
        this.Controls.Add(restartButton);

        var selectedImages = imagePaths.OrderBy(x => random.Next()).Take(3).ToArray();

        for (int i = 0; i < selectedImages.Length; i++)
        {
            var imagePath = selectedImages[i];
            var png = Path.Combine(homePath, "Projects\\2024\\rockyourslotsoff\\NETSlot\\Reel-Images", imagePath);
            Console.WriteLine(png);
            try
            {
                PictureBox pictureBox = new PictureBox
                {
                    Image = Image.FromFile(png),
                    SizeMode = PictureBoxSizeMode.Zoom,
                    Width = 256,
                    Height = 256,
                    Margin = new Padding(0),
                    Location = new Point(i * 256 + 10, this.Height / 4 - 256 / 2)
                };
                this.Controls.Add(pictureBox);
            }
            catch (FileNotFoundException)
            {
                Console.WriteLine($"File not found: {png}");
            }
            catch (Exception ex)
            {
                Console.WriteLine($"Error loading image: {ex.Message}");
            }
        }
    }

    private void DurationTimer_Tick(object? sender, EventArgs e)
    {
        timer.Stop();
        durationTimer.Stop();
    }

    private void RestartButton_Click(object? sender, EventArgs e)
    {
        StartTimers();
    }

    private void StartTimers()
    {
        timer.Start();
        durationTimer.Start();
    }

    [STAThread]
    public static void Main()
    {
        Application.EnableVisualStyles();
        Application.SetCompatibleTextRenderingDefault(false);
        Application.Run(new MainForm());
    }
}